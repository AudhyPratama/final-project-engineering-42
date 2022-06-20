import React, { useState } from 'react'


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
    ]);


    return (
        <>
            <h1>Daftar Buku</h1>
            <div className="products">
                {products.map((products, idx) => (
                    <div className="product" key={idx}>
                        <h3>{products.name}</h3>
                        <h4>{products.cost}</h4>
                        <img src={products.image} alt={products.name} />
                        <button onClick={() => addToCart(products)}>+</button>
                    </div>
                ))}
            </div>
        </>
    );
}

export default Products