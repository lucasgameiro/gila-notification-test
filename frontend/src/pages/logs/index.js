import styles from '@/styles/Logs.module.css'
import Layout from '../../components/layout'
import { loadLogs } from '../../lib/load-logs'

export async function getServerSideProps() {
  const logs = await loadLogs()
  const log = logs.join('\n')
  return { props: { log } }
}

export default function Logs({ log }) {
  return (
    <Layout>
      <div className={styles.description}>
        <div>
          <h1>Notification test</h1>
          <p>List of message sent in this instance</p>
        </div>
      </div>
      <div className={styles.center}>
        <textarea value={log} readOnly></textarea>
      </div>
    </Layout>
  )
}
