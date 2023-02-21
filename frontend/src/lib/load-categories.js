export async function loadCategories() {
  const res = await fetch(process.env.BACKEND_ADDRESS + '/categories/')
  const data = await res.json()
  return data
}
