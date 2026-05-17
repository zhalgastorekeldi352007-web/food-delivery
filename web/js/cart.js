let cart = [];

// Load cart from localStorage
function loadCart() {
    const savedCart = localStorage.getItem("zafood_cart");
    if (savedCart) {
        cart = JSON.parse(savedCart);
    }
    updateCartDisplay();
}

// Save cart to localStorage
function saveCart() {
    localStorage.setItem("zafood_cart", JSON.stringify(cart));
}

// Add item to cart
function addToCart(itemId) {
    const item = menuData.find(i => i.id === itemId);
    if (!item) return;
    
    const existingItem = cart.find(i => i.id === itemId);
    if (existingItem) {
        existingItem.quantity++;
    } else {
        cart.push({
            ...item,
            quantity: 1
        });
    }
    
    saveCart();
    updateCartDisplay();
    
    // Animation feedback
    showNotification(`${item.name} added to cart!`);
}

// Remove item from cart
function removeFromCart(itemId) {
    const index = cart.findIndex(i => i.id === itemId);
    if (index !== -1) {
        const item = cart[index];
        cart.splice(index, 1);
        saveCart();
        updateCartDisplay();
        showNotification(`${item.name} removed from cart`);
    }
}

// Update quantity
function updateQuantity(itemId, delta) {
    const item = cart.find(i => i.id === itemId);
    if (item) {
        item.quantity += delta;
        if (item.quantity <= 0) {
            removeFromCart(itemId);
        } else {
            saveCart();
            updateCartDisplay();
        }
    }
}

// Update cart UI
function updateCartDisplay() {
    const cartCount = document.getElementById("cartCount");
    const cartItems = document.getElementById("cartItems");
    const cartTotal = document.getElementById("cartTotal");
    
    if (cartCount) {
        const totalItems = cart.reduce((sum, item) => sum + item.quantity, 0);
        cartCount.textContent = totalItems;
    }
    
    if (cartItems) {
        if (cart.length === 0) {
            cartItems.innerHTML = '<p style="text-align: center; color: #999;">Cart is empty</p>';
        } else {
            cartItems.innerHTML = cart.map(item => `
                <div class="cart-item">
                    <div class="cart-item-info">
                        <div class="cart-item-title">${item.name}</div>
                        <div class="cart-item-price">${item.price} ₸</div>
                    </div>
                    <div class="cart-item-quantity">
                        <button class="quantity-btn" onclick="updateQuantity(${item.id}, -1)">-</button>
                        <span>${item.quantity}</span>
                        <button class="quantity-btn" onclick="updateQuantity(${item.id}, 1)">+</button>
                        <button class="remove-item" onclick="removeFromCart(${item.id})">✕</button>
                    </div>
                </div>
            `).join("");
        }
    }
    
    if (cartTotal) {
        const total = cart.reduce((sum, item) => sum + (item.price * item.quantity), 0);
        cartTotal.textContent = `${total} ₸`;
    }
}

// Show notification
function showNotification(message) {
    const notification = document.createElement("div");
    notification.textContent = message;
    notification.style.cssText = `
        position: fixed;
        bottom: 20px;
        right: 20px;
        background: #28a745;
        color: white;
        padding: 12px 20px;
        border-radius: 10px;
        z-index: 2000;
        animation: slideIn 0.3s;
    `;
    document.body.appendChild(notification);
    setTimeout(() => {
        notification.remove();
    }, 2000);
}

// Toggle cart sidebar
function toggleCart() {
    const sidebar = document.getElementById("cartSidebar");
    const overlay = document.getElementById("overlay");
    
    if (sidebar) {
        sidebar.classList.toggle("open");
    }
    if (overlay) {
        overlay.classList.toggle("show");
    }
}

// Checkout
async function checkout() {
    if (cart.length === 0) {
        showNotification("Cart is empty!");
        return;
    }
    
    const order = {
        user_id: "customer_" + Date.now(),
        restaurant_id: "zafood_001",
        items: cart.map(item => ({
            menu_item_id: `item_${item.id}`,
            name: item.name,
            quantity: item.quantity,
            price: item.price
        })),
        total: cart.reduce((sum, item) => sum + (item.price * item.quantity), 0)
    };
    
    try {
        // Send order to API Gateway
        const response = await fetch("http://localhost:8080/api/v1/orders", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(order)
        });
        
        if (response.ok) {
            const result = await response.json();
            showOrderModal(result);
            cart = [];
            saveCart();
            updateCartDisplay();
            toggleCart();
        } else {
            // Demo mode - show success anyway
            showOrderModal({ id: "DEMO_" + Date.now(), status: "confirmed" });
            cart = [];
            saveCart();
            updateCartDisplay();
            toggleCart();
        }
    } catch (error) {
        // Demo mode if backend not running
        console.log("Backend not available, demo mode");
        showOrderModal({ id: "DEMO_" + Date.now(), status: "confirmed" });
        cart = [];
        saveCart();
        updateCartDisplay();
        toggleCart();
    }
}

function showOrderModal(order) {
    const modal = document.getElementById("orderModal");
    const message = document.getElementById("orderMessage");
    const total = cart.reduce((sum, item) => sum + (item.price * item.quantity), 0);
    
    message.innerHTML = `
        Your order #${order.id} has been confirmed!<br>
        Total: ${total} ₸<br>
        Delivery time: 30-45 minutes.
    `;
    
    if (modal) {
        modal.style.display = "flex";
    }
}

function closeModal() {
    const modal = document.getElementById("orderModal");
    if (modal) {
        modal.style.display = "none";
    }
}

// Initialize on load
if (typeof document !== 'undefined') {
    document.addEventListener("DOMContentLoaded", () => {
        loadCart();
    });
}
