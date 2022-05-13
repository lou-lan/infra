import useSWR, { useSWRConfig } from 'swr'
import Head from 'next/head'
import { useState, useMemo } from 'react'
import { useTable } from 'react-table'
import dayjs from 'dayjs'
import { ShareIcon, XIcon } from '@heroicons/react/outline'

import { useAdmin } from '../../lib/admin'

import Dashboard from '../../components/layouts/dashboard'
import Loader from '../../components/loader'
import Table from '../../components/table'
import EmptyTable from '../../components/empty-table'
import DeleteModal from '../../components/modals/delete'
import Grant from '../../components/grant'
import PageHeader from '../../components/layouts/page-header'

function columns (admin) {
  return [
    {
      Header: 'Cluster',
      accessor: 'name',
      Cell: ({ value }) => (
        <div className='flex items-center'>
          <div className='py-2'>{value}</div>
        </div>
      )
    }, {
      Header: 'Added',
      accessor: i => {
        return dayjs(i.created).fromNow()
      }
    },
    ...admin
      ? [{
          id: 'access',
          accessor: i => i,
          Header: () => (
            <div className='text-right'>
              Access
            </div>
          ),
          Cell: ({ row }) => {
            const [shareOpen, setShareOpen] = useState(false)
            const { data: grants } = useSWR(`/v1/grants?resource=${row.original.name}`)

            const users = new Set(grants?.filter(g => !g?.subject?.startsWith('g:'))).size
            const groups = new Set(grants?.filter(g => g?.subject?.startsWith('g:'))).size

            return (
              <div className='flex text-right justify-end w-24 h-8 ml-auto'>
                {grants && (
                  <div className='group-hover:hidden flex justify-center items-center text-gray-300'>
                    {users === 0 && groups === 0
                      ? (
                        <div>
                          No access
                        </div>
                        )
                      : (
                        <>
                          {users > 0 && (
                            <div>
                              {users}&nbsp;User{users > 1 && 's'}
                            </div>
                          )}
                          {users > 0 && groups > 0 && (
                            <div className='mx-1'>•</div>
                          )}
                          {groups > 0 && (
                            <div>
                              {groups}&nbsp;Group{groups > 1 && 's'}
                            </div>
                          )}
                        </>
                        )}
                  </div>
                )}
                <div className='group-hover:flex space-x-1 hidden'>
                  <button onClick={() => setShareOpen(true)} className='cursor-pointer bg-zinc-900 rounded-lg'>
                    <div className='flex items-center py-1 px-3 text-gray-300 hover:text-white'>
                      <ShareIcon className='w-4 h-4 ' /><div className='text-sm ml-1'>Share</div>
                    </div>
                  </button>

                  {/* grant modal */}
                  <Grant
                    id={row.original.id}
                    modalOpen={shareOpen}
                    handleCloseModal={() => setShareOpen(false)}
                  />
                </div>
              </div>
            )
          }
        }]
      : [],
    ...admin
      ? [{
          id: 'remove',
          accessor: d => d,
          Cell: ({ rows, value: { id, name } }) => {
            const { mutate } = useSWRConfig()

            const [open, setOpen] = useState(false)

            return (
              <div className='flex justify-end w-6 ml-auto opacity-0 group-hover:opacity-100'>
                <button onClick={() => setOpen(true)} className='py-1 px-2 -mr-2 cursor-pointer'>
                  <XIcon className='w-5 h-5 text-gray-500 hover:text-white' />
                </button>

                {/* delete modal */}
                <DeleteModal
                  open={open}
                  setOpen={setOpen}
                  onSubmit={async () => {
                    mutate('/v1/destinations', async destinations => {
                      await fetch(`/v1/destinations/${id}`, {
                        method: 'DELETE'
                      })

                      return destinations?.filter(d => d?.id !== id)
                    }, { optimisticData: rows.map(r => r.original).filter(d => d?.id !== id) })

                    setOpen(false)
                  }}
                  title='Delete Cluster'
                  message={<>Are you sure you want to disconnect <span className='text-white font-bold'>{name}?</span><br />Note: you must also uninstall the Infra Connector from this cluster.</>}
                />
              </div>
            )
          }
        }]
      : []
  ]
}

export default function Destinations () {
  const { data: destinations, error } = useSWR('/v1/destinations')
  const { admin, loading: adminLoading } = useAdmin()
  const table = useTable({ columns: useMemo(() => columns(admin), [admin]), data: destinations || [] })

  const loading = adminLoading || (!destinations && !error)

  return (
    <>
      <Head>
        <title>Destinations - Infra</title>
      </Head>
      {loading
        ? (<Loader />)
        : (
          <div className='flex-1 flex flex-col space-y-8 mt-3 mb-4'>
            <PageHeader header='Infrastructure' buttonHref={admin && '/destinations/add'} buttonLabel='Infrastructure' />
            {error?.status
              ? <div className='my-20 text-center font-light text-gray-300 text-sm'>{error?.info?.message}</div>
              : <>
                <Table {...table} />
                {
                    destinations?.length === 0 &&
                      <EmptyTable
                        title='There are no infrastructure'
                        subtitle={`There are currently no infrastructure connected to Infra. ${admin ? 'Get started by connecting one.' : ''}`}
                        iconPath='/destinations.svg'
                        buttonHref={admin && '/destinations/add'}
                        buttonText='Infrastructure'
                      />
                  }
                </>}
          </div>
          )}
    </>
  )
}

Destinations.layout = function (page) {
  return (
    <Dashboard>
      {page}
    </Dashboard>
  )
}
