export async function loadLogs() {
  const res = await fetch(process.env.BACKEND_ADDRESS + '/log/')
  const data = await res.json()
  return data
}
