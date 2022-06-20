import React from 'react'

function Cart({cart, removeFromCart}) {
    return (
        <>
            <h1>Daftar Pesanan</h1>
            <h5>Jumlah Pesanan : {cart.length} item</h5>
            <div className="products">
                {cart.map((products, idx) => (
                    <div className="product" key={idx}>
                        <h3>{products.name}</h3>
                        <h4>{products.cost}</h4>
                        <img src={products.image} alt={products.name} />
                        <button onClick={() => removeFromCart(products)}>Hapus</button>
                    </div>
                ))}
            </div>
        </>
    )
}

export default Cart