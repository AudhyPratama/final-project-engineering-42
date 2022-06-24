import React from 'react'
import { Button } from 'react-bootstrap'

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
                        <Button variant="danger" onClick={() => removeFromCart(products)}>Hapus</Button>
                    </div>
                ))}
            </div>
        </>
    )
}

export default Cart