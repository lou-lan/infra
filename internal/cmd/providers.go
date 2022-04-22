package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/infrahq/infra/api"
	"github.com/infrahq/infra/internal/cmd/cliopts"
)

func newProvidersCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "providers",
		Short:   "Manage identity providers",
		Aliases: []string{"provider"},
		Group:   "Management commands:",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			if err := rootPreRun(cmd.Flags()); err != nil {
				return err
			}
			return mustBeLoggedIn()
		},
	}

	cmd.AddCommand(newProvidersListCmd())
	cmd.AddCommand(newProvidersAddCmd())
	cmd.AddCommand(newProvidersRemoveCmd())

	return cmd
}

func newProvidersListCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List connected identity providers",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := defaultAPIClient()
			if err != nil {
				return err
			}

			providers, err := client.ListProviders("")
			if err != nil {
				return err
			}

			type row struct {
				Name string `header:"NAME"`
				URL  string `header:"URL"`
			}

			var rows []row
			for _, p := range providers {
				rows = append(rows, row{Name: p.Name, URL: p.URL})
			}

			if len(rows) > 0 {
				printTable(rows)
			} else {
				fmt.Println("No providers found")
			}

			return nil
		},
	}
}

type providerAddOptions struct {
	URL          string
	ClientID     string
	ClientSecret string
}

func (o providerAddOptions) Validate() error {
	var missing []string
	if o.URL == "" {
		missing = append(missing, "url")
	}
	if o.ClientID == "" {
		missing = append(missing, "client-id")
	}
	if o.ClientSecret == "" {
		missing = append(missing, "client-secret")
	}
	if len(missing) > 0 {
		return fmt.Errorf("missing value for required flags: %v", strings.Join(missing, ", "))
	}
	return nil
}

func newProvidersAddCmd() *cobra.Command {
	var opts providerAddOptions

	cmd := &cobra.Command{
		Use:   "add PROVIDER",
		Short: "Connect an identity provider",
		Long: `
Add an identity provider for users to authenticate.

PROVIDER is a short unique name of the identity provider bieng added (eg. okta) 
		`,
		Args: ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cliopts.DefaultsFromEnv("INFRA_PROVIDER", cmd.Flags()); err != nil {
				return err
			}

			if err := opts.Validate(); err != nil {
				return err
			}

			client, err := defaultAPIClient()
			if err != nil {
				return err
			}

			_, err = client.CreateProvider(&api.CreateProviderRequest{
				Name:         args[0],
				URL:          opts.URL,
				ClientID:     opts.ClientID,
				ClientSecret: opts.ClientSecret,
			})
			if err != nil {
				return err
			}

			fmt.Printf("Provider %s added\n", args[0])
			return nil
		},
	}

	cmd.Flags().StringVar(&opts.URL, "url", "", "Base URL of the domain of the OIDC identity provider (eg. acme.okta.com)")
	cmd.Flags().StringVar(&opts.ClientID, "client-id", "", "OIDC client ID")
	cmd.Flags().StringVar(&opts.ClientSecret, "client-secret", "", "OIDC client secret")
	return cmd
}

func newProvidersRemoveCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "remove PROVIDER",
		Aliases: []string{"rm"},
		Short:   "Disconnect an identity provider",
		Args:    ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := defaultAPIClient()
			if err != nil {
				return err
			}

			providerName := args[0]

			providers, err := client.ListProviders(providerName)
			if err != nil {
				return err
			}

			switch len(providers) {
			case 0:
				return fmt.Errorf("Cannot remove provider %s: not found", providerName)
			case 1:
				if err := client.DeleteProvider(providers[0].ID); err != nil {
					return err
				}

				fmt.Fprintf(os.Stderr, "Provider %s removed\n", providerName)
			default:
				panic(fmt.Sprintf(DuplicateEntryPanic, "provider", providerName))
			}

			return nil
		},
	}
}

func GetProviderByName(client *api.Client, name string) (*api.Provider, error) {
	providers, err := client.ListProviders(name)
	if err != nil {
		return nil, err
	}

	if len(providers) == 0 {
		return nil, fmt.Errorf("no identity providers connected with the name %s", name)
	}

	return &providers[0], nil
}
