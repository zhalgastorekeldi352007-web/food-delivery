const menuData = [
    // Pizza
    {
        id: 1,
        name: "Margherita",
        description: "Tomato sauce, mozzarella, basil",
        price: 2500,
        category: "pizza",
        image: "🍕",
        emoji: "🍕"
    },
    {
        id: 2,
        name: "Pepperoni",
        description: "Pepperoni, mozzarella, tomato sauce",
        price: 3200,
        category: "pizza",
        image: "🍕",
        emoji: "🍕"
    },
    {
        id: 3,
        name: "Hawaiian",
        description: "Chicken, pineapple, mozzarella",
        price: 3000,
        category: "pizza",
        image: "🍕",
        emoji: "🍍"
    },
    {
        id: 4,
        name: "Four Cheese",
        description: "Mozzarella, parmesan, gorgonzola, feta",
        price: 3500,
        category: "pizza",
        image: "🧀",
        emoji: "🧀"
    },
    // Burgers
    {
        id: 5,
        name: "Classic Burger",
        description: "Beef, lettuce, tomato, cheese",
        price: 1800,
        category: "burger",
        image: "🍔",
        emoji: "🍔"
    },
    {
        id: 6,
        name: "Cheeseburger",
        description: "Double cheese, beef, special sauce",
        price: 2200,
        category: "burger",
        image: "🍔",
        emoji: "🧀"
    },
    {
        id: 7,
        name: "Spicy Burger",
        description: "Jalapeno, spicy sauce, beef",
        price: 2400,
        category: "burger",
        image: "🌶️",
        emoji: "🌶️"
    },
    // Sushi
    {
        id: 8,
        name: "Philadelphia Roll",
        description: "Salmon, cream cheese, cucumber",
        price: 4500,
        category: "sushi",
        image: "🍣",
        emoji: "🍣"
    },
    {
        id: 9,
        name: "California Roll",
        description: "Crab, avocado, cucumber, tobiko",
        price: 4200,
        category: "sushi",
        image: "🍣",
        emoji: "🦀"
    },
    {
        id: 10,
        name: "Roll Set",
        description: "Assorted 6 types of rolls",
        price: 8900,
        category: "sushi",
        image: "🍱",
        emoji: "🍱"
    },
    // Drinks
    {
        id: 11,
        name: "Coca-Cola",
        description: "0.5L",
        price: 500,
        category: "drinks",
        image: "🥤",
        emoji: "🥤"
    },
    {
        id: 12,
        name: "Orange Juice",
        description: "Fresh, 0.3L",
        price: 600,
        category: "drinks",
        image: "🧃",
        emoji: "🧃"
    },
    {
        id: 13,
        name: "Cappuccino",
        description: "Italian coffee",
        price: 700,
        category: "drinks",
        image: "☕",
        emoji: "☕"
    },
    {
        id: 14,
        name: "Milkshake",
        description: "Vanilla/Chocolate/Strawberry",
        price: 800,
        category: "drinks",
        image: "🥛",
        emoji: "🥛"
    }
];

let currentFilter = "all";

function displayMenu() {
    const menuGrid = document.getElementById("menuGrid");
    if (!menuGrid) return;
    
    const filteredItems = currentFilter === "all" 
        ? menuData 
        : menuData.filter(item => item.category === currentFilter);
    
    menuGrid.innerHTML = filteredItems.map(item => `
        <div class="menu-item">
            <div class="menu-item-image">
                ${item.image || item.emoji}
            </div>
            <div class="menu-item-info">
                <h3>${item.name}</h3>
                <p>${item.description}</p>
                <div class="price">${item.price} ₸</div>
                <button class="add-to-cart" onclick="addToCart(${item.id})">
                    🛒 Add to Cart
                </button>
            </div>
        </div>
    `).join("");
}

function filterCategory(category) {
    currentFilter = category;
    displayMenu();
    
    // Update active button
    document.querySelectorAll(".filter-btn").forEach(btn => {
        btn.classList.remove("active");
        if (btn.textContent.toLowerCase().includes(category) || 
            (category === "all" && btn.textContent === "All")) {
            btn.classList.add("active");
        }
    });
}

// Load menu on page load
if (typeof document !== 'undefined') {
    document.addEventListener("DOMContentLoaded", displayMenu);
}
