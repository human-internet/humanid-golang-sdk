import Head from 'next/head'

export default function Home() {
  return (
    <div>
      <Head>
        <title>HumanID Demo</title>
        <link rel="icon" href="/download.png" />
      </Head>

      <main>
        <a href="/request"><img src="humanid-login.png" alt="Anonymous Login with humanID" height="60" /></a>
      </main>
    </div>
  )
}
