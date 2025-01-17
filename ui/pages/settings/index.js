import Head from 'next/head'

import Dashboard from '../../components/layouts/dashboard'
import HeaderIcon from '../../components/header-icon'
import Admin from './admin'

export default function Settings () {
  return (
    <>
      <Head>
        <title>Settings - Infra</title>
      </Head>
      <div className='flex flex-row mt-4 mb-4 lg:mt-6'>
        <div className='mt-2 mr-6'>
          <HeaderIcon iconPath='/settings-color.svg' />
        </div>
        <div className='flex-1 flex flex-col space-y-4'>
          <h1 className='text-base font-bold mt-6 mb-4'>Settings</h1>
          <div className='pt-3'>
            <Admin />
          </div>
        </div>
      </div>
    </>
  )
}

Settings.layout = function (page) {
  return (
    <Dashboard>
      {page}
    </Dashboard>
  )
}
