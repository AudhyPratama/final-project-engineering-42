import React from 'react'
import { Button, CardGroup, Card} from 'react-bootstrap'
import { useState } from 'react'
import axios from 'axios'
import { useNavigate } from 'react-router-dom'

function Cart({cart, removeFromCart}) {
    const [cartItems, setCartItems] = useState(cart)
    const [total, setTotal] = useState(0)

    const removeItem = (id) => {
        setCartItems(cartItems.filter(item => item.id !== id))
    }

    const getCart = () => {
        axios.get('http://localhost:8080/api/carts')
        .then(res => {
            setCartItems(res.data)
        }
        ).catch(err => {
            console.log(err)
        });
    };
    
    const removeAll = () => {
        axios.get('http://localhost:8080/api/cart/delete-all')
        .then(res => {
            setCartItems([])
        }
        ).catch(err => {
            console.log(err)
        });
    };
    
    const removeItemFromCart = (id) => {
        axios.post('http://localhost:8080/api/cart/delete/' + id)
        .then(res => {
            setCartItems(cartItems.filter(item => item.id !== id))
        }
        ).catch(err => {
            console.log(err)
        });
    };

    
    const removeAllFromCart = () => {
        setCartItems([])
    }

    const navigate = useNavigate();

    return (
        <>
            <h1>Daftar Pesanan</h1>
            <h5>Jumlah Pesanan : {cartItems.length}</h5>
            <table className="table table-striped">
                <thead>
                    <tr>
                        <th>Gambar</th>
                        <th>Judul</th>
                        <th>Kategori</th>
                        <th>Harga</th>
                        <th>Jumlah</th>
                        <th>Subtotal</th>
                        <th>Aksi</th>
                    </tr>
                </thead>
                <tbody>
                    {cartItems.map(item => (
                        <tr key={item.id}>
                            <td><img src={item.image} alt={item.name} width="100" /></td>
                            <td>{item.name}</td>
                            <td>{item.category_name}</td>
                            <td>Rp. {item.harga}</td>
                            <td>
                                <input type="number" value={item.quantity} onChange={(e) => {
                                    const newCart = cartItems.map(cartItem => {
                                        if (cartItem.id === item.id) {
                                            cartItem.quantity = e.target.value
                                        }
                                        return cartItem
                                    })
                                    setCartItems(newCart)
                                }} />
                            </td>
                            <td>Rp. {item.harga * item.quantity}</td>
                            <td>
                                <Button variant="danger" onClick={() =>removeItem(item.id)}>Hapus</Button>
                            </td>
                        </tr>

                    ))} 
                </tbody>
            </table>
            <h2>Total : Rp.
                {cartItems.reduce((total, item) => {
                    return total + (item.harga * item.quantity)
                }, 0)}
            </h2>
            <Button variant="danger" onClick={removeAllFromCart}>Reset</Button>
            <Button variant="success" onClick={() => navigate('/payment')}>Checkout</Button>
        </>
    )
}

export default Cart