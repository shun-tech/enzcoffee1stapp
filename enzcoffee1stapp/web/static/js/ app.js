document.addEventListener('DOMContentLoaded', function() {
    fetchProducts();
});

function fetchProducts() {
    fetch('/api/products')
        .then(response => response.json())
        .then(data => {
            const productList = document.getElementById('product-list');
            productList.innerHTML = '';
            data.forEach(product => {
                productList.innerHTML += `<div>${product.name} - ${product.price}円</div>`;
            });
        })
        .catch(error => console.error('Error fetching products:', error));
}

function addProduct() {
    const name = document.getElementById('productName').value;
    const price = document.getElementById('productPrice').value;
    fetch('/api/products', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ name: name, price: parseFloat(price) })
    })
    .then(response => response.json())
    .then(data => {
        console.log('Product added:', data);
        fetchProducts();  // 更新された商品リストを再取得
    })
    .catch(error => console.error('Error adding product:', error));
}