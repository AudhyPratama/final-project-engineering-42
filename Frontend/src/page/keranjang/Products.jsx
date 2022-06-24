import React, { useState } from 'react'
import { CardGroup, Card, Button, Container } from 'react-bootstrap'
import { BsFillCartPlusFill } from "react-icons/bs";


function Products({addToCart}) {

    const [products] = useState([
        {
            name: 'Practical MongoDB',
            cost: 'Rp.78.000',
            image: 'https://itbook.store/img/books/9781484206485.png'
        },
        {
            name: 'Designing Across Senses',
            cost: 'Rp.72.000',
            image: 'https://itbook.store/img/books/9781491954249.png'
        },
        {
            name: 'Designing Across Senses',
            cost: 'Rp.72.000',
            image: 'https://itbook.store/img/books/9781491954249.png'
        },
        {
            name: 'Designing Across Senses',
            cost: 'Rp.72.000',
            image: 'https://itbook.store/img/books/9781491954249.png'
        },

        {
            name: 'Designing Across Senses',
            cost: 'Rp.72.000',
            image: 'https://itbook.store/img/books/9781491954249.png'
        },

        {
            name: 'Designing Across Senses',
            cost: 'Rp.72.000',
            image: 'https://itbook.store/img/books/9781491954249.png'
        },
    ]);


    return (
        <>
            <h1>Daftar Buku</h1>
            <div className="products">
                {products.map((products, idx) => (
                    <div className="product" key={idx}>

                            <CardGroup style={{ width: "14rem" }} CardGroup >
                                <Card >
                                    <Card.Img variant="top" src={products.image} />
                                    <Card.Body>
                                        <Card.Title>{products.name}</Card.Title>
                                        <h5>{products.cost}</h5>
                                        <Card.Text>
                                            This is a wider card with supporting text below as a natural lead-in to
                                            additional content.
                                        </Card.Text>
                                    </Card.Body>
                                    <Card.Footer>
                                        <Button variant="outline-success" onClick={() => addToCart(products)}><BsFillCartPlusFill /> Keranjang</Button>
                                    </Card.Footer>
                                </Card>
                            </CardGroup>

                            <br></br>
                        {/* <h3>{products.name}</h3> */}
                        {/* <h4>{products.cost}</h4> */}
                        {/* <img src={products.image} alt={products.name} /> */}
                        {/* <button onClick={() => addToCart(products)}>Tambah Ke Keranjang</button> */}
                    </div>
                ))}
            </div>
        </>
    );
}

export default Products