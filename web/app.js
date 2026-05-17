const apiBase = 'http://localhost:8080/api'
let token = ''

async function request(path, method='GET', body) {
  const opts = { method, headers: { 'Content-Type': 'application/json' } }
  if (token) opts.headers.Authorization = token
  if (body) opts.body = JSON.stringify(body)
  const res = await fetch(`${apiBase}${path}`, opts)
  return res.json()
}

async function registerUser() {
  const email = document.getElementById('email').value
  const password = document.getElementById('password').value
  const name = document.getElementById('name').value
  const data = await request('/register', 'POST', { email, password, name })
  if (data.token) { token = data.token; alert('Registered') }
}

async function loginUser() {
  const email = document.getElementById('email').value
  const password = document.getElementById('password').value
  const data = await request('/login', 'POST', { email, password })
  if (data.token) { token = data.token; alert('Logged in') }
}

async function loadRestaurants() {
  const list = document.getElementById('restaurantList')
  const data = await request('/restaurants')
  list.innerHTML = JSON.stringify(data, null, 2)
}

async function createOrder() {
  const restaurantId = document.getElementById('restaurantId').value
  const item = {
    menu_item_id: document.getElementById('itemId').value,
    name: document.getElementById('itemName').value,
    quantity: parseInt(document.getElementById('itemQty').value, 10),
    price: parseFloat(document.getElementById('itemPrice').value)
  }
  const data = await request('/orders', 'POST', { restaurant_id: restaurantId, items: [item] })
  document.getElementById('orderResult').innerText = JSON.stringify(data, null, 2)
}

async function loadOrder() {
  const id = document.getElementById('orderId').value
  const data = await request(`/orders/${id}`)
  document.getElementById('orderDetails').innerText = JSON.stringify(data, null, 2)
}

async function addMenuItem() {
  const name = document.getElementById('menuName').value
  const price = parseFloat(document.getElementById('menuPrice').value)
  const data = await request('/menu', 'POST', { name, price })
  alert(JSON.stringify(data))
}

async function assignDelivery() {
  const orderId = document.getElementById('deliverOrderId').value
  const personId = document.getElementById('deliveryPersonId').value
  const data = await request('/deliveries/assign', 'POST', { order_id: orderId, delivery_person_id: personId })
  alert(JSON.stringify(data))
}
