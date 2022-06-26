import axios from "axios";
import React, { useState } from "react";
import { useEffect } from "react";
import { CardGroup, Card, Button, Container } from "react-bootstrap";
import { BsFillCartPlusFill } from "react-icons/bs";

function Products({ addToCart }) {
  const [products, setProducts] = useState([]);
  const fetchProducts = async () => {
    const resp = await axios.get("http://localhost:8080/api/products");
    const data = resp.data;
    // console.log(data);
    setProducts(data.products);
  };

  useEffect(() => {
    fetchProducts()
      //   .then(v =>)
      .catch((err) => console.error(err));
  }, []);

  return (
    <>
      <h1>Daftar Buku</h1>
      <div className="products">
        {products.map((products, idx) => (
          <div className="product" key={idx}>
            <CardGroup style={{ width: "20rem" }} CardGroup>
              <Card>
                <Card.Img variant="top" src={products.image} />
                <Card.Body>
                  <Card.Title>{products.name}</Card.Title>
                  <h5>Rp. {products.harga}</h5>
                  <Card.Text>
                  <h6>Kondisi : {products.kondisi} </h6>
                  <h6>Stok : {products.stock} </h6>
                  <h6>Berat : {products.berat} </h6>
                  <h6>Penulis : {products.penulis} </h6>
                  <h6>Penerbit : {products.penerbit} </h6>
                  <h6>Kategori : {products.category_name} </h6>
                  <h7>{products.deskripsi} </h7>
                  </Card.Text>
                </Card.Body>
                <Card.Footer>
                  <Button
                    variant="outline-success"
                    onClick={() => addToCart(products)}
                  >
                    <BsFillCartPlusFill /> Keranjang
                  </Button>
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

export default Products;
