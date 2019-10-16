import React from 'react'
import Head from 'next/head'
import { SettingsTab } from '../../src/user/interfaces'
import Settings from '../../src/user/pages/Settings'
import { useRouter } from 'next/router'

const SettingsPage = () => {
  const router = useRouter()
  const currentTab = router.query.tab as SettingsTab

  return (
    <>
      <Head>
        <title>Tadoku - Settings</title>
      </Head>
      <Settings tab={currentTab} />
    </>
  )
}

export default SettingsPage
