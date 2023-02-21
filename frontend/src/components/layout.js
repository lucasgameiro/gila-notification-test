import Head from 'next/head'
import Link from 'next/link'
import { Inter } from '@next/font/google'
import styles from '@/styles/Layout.module.css'
import { useRouter } from 'next/router'

const inter = Inter({ subsets: ['latin'] })

export default function Layout({ children }) {
  const router = useRouter()
  return (
    <>
      <Head>
        <title>Notification Test</title>
        <meta
          name="description"
          content="Notification test for Gila Software"
        />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <main className={styles.main}>
        {children}
        <div className={styles.grid}>
          <Link
            href="/"
            className={`${styles.card} ${
              router.pathname == '/' ? styles.cardActive : ''
            }`}
            rel="noopener noreferrer"
          >
            <h2 className={inter.className}>
              Post Message <span>-&gt;</span>
            </h2>
            <p className={inter.className}>
              Post new message to category channel.
            </p>
          </Link>
          <Link
            href="/logs"
            className={`${styles.card} ${
              router.pathname == '/logs' ? styles.cardActive : ''
            }`}
            rel="noopener noreferrer"
          >
            <h2 className={inter.className}>
              Logs <span>-&gt;</span>
            </h2>
            <p className={inter.className}>
              Check messages sent in this instance.
            </p>
          </Link>
        </div>
      </main>
    </>
  )
}
