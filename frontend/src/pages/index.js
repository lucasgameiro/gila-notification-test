import styles from '@/styles/Home.module.css'
import { loadCategories } from '../lib/load-categories'
import Layout from '../components/layout'
import { useState } from 'react'

export async function getServerSideProps() {
  const categories = await loadCategories()

  return { props: { categories } }
}

export default function Home({ categories }) {
  const [categoryId, setCategoryId] = useState(0)
  const [message, setMessage] = useState('')
  const [feedback, setFeedback] = useState('')
  const sendMessage = () => {
    fetch('/api/message', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded',
      },
      body:
        'category_id=' +
        categoryId +
        '&message=' +
        window.encodeURIComponent(message),
    })
      .then((response) => {
        if (response.ok) {
          setCategoryId(0)
          setMessage('')
          setFeedback('Message sent...')
          window.setTimeout(() => setFeedback(''), 2000)
        } else {
          setFeedback('Message could not be sent. Try again later...')
        }
      })
      .catch((e) => console.error(e))
  }

  return (
    <Layout>
      <div className={styles.description}>
        <div>
          <h1>Notification test</h1>
          <p>
            Choose category, write text message and click Submit to send
            notification to subscribed users
          </p>
        </div>
      </div>
      <div className={styles.center}>
        <form
          onSubmit={(e) => {
            e.preventDefault()
            sendMessage()
          }}
        >
          <div className={styles.formRow}>
            <label htmlFor="category">Category</label>
            <select
              required
              id="category"
              name="category_id"
              value={categoryId}
              onChange={(e) => setCategoryId(e.target.value)}
            >
              <option></option>
              {categories.map((category) => (
                <option value={category.id} key={category.id}>
                  {category.name}
                </option>
              ))}
            </select>
          </div>
          <div className={styles.formRow}>
            <label htmlFor="message">Message</label>
            <input
              required
              type="text"
              id="message"
              name="message"
              value={message}
              onChange={(e) => setMessage(e.target.value)}
            />
          </div>
          <div className={styles.formRow}>
            <button type="submit">Submit</button>
          </div>
          <div className={styles.formRow}>{feedback}</div>
        </form>
      </div>
    </Layout>
  )
}
