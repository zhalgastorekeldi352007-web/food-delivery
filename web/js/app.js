// Main app initialization
console.log("ZAfood App Loaded!");

// Auto-refresh cart count
setInterval(() => {
    if (typeof updateCartDisplay === 'function') {
        updateCartDisplay();
    }
}, 1000);

// Close modal on outside click
window.onclick = function(event) {
    const modal = document.getElementById("orderModal");
    if (event.target === modal) {
        closeModal();
    }
};
