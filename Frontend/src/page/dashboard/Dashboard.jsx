// import module
import React, { useState } from "react";
import "bootstrap/dist/css/bootstrap.min.css";
import {
    Navbar,
    Container,
    Nav,
    Button,
    Form,
    FormControl,
} from "react-bootstrap";
// import DashboardUi from './page/dashboard/Dashboard';

// import react icon
import { BsCartFill } from "react-icons/bs";
import { BsFillPersonFill } from "react-icons/bs";
import HomeCarousel from "../keranjang/HomeCarousel";

// import file
import Products from "../keranjang/Products";
import Cart from "../keranjang/Cart";

// import css
import "./Dashboard.css";

const Dashboard = () => {
    // keranjang belanja
    const PAGE_PRODUCTS = "products";
    const PAGE_CART = "cart";

    const [cart, setCart] = useState([PAGE_CART]);
    const [page, setPage] = useState([PAGE_PRODUCTS]);

    const removeFromCart = (productToRemove) => {
        setCart(cart.filter((product) => product !== productToRemove));
    };

    const addToCart = (product) => {
        console.log("add cart");
        setCart([...cart, { ...product }]);
    };

    const navigateTo = (nextPage) => {
        setPage(nextPage);
    };
    return (
        <div className="Dashboard">
            {/* navbar */}
            <Navbar bg="light" expand="lg">
                <Container>
                    <Navbar.Brand href="#home">PesBuk</Navbar.Brand>
                    <Navbar.Toggle aria-controls="basic-navbar-nav" />
                    <Navbar.Collapse id="basic-navbar-nav">
                        <Nav className="me-auto">
                            <Nav.Link href="#home">Home</Nav.Link>
                            <Nav.Link href="#link">Link</Nav.Link>
                        </Nav>
                        <Form className="d-flex">
                            <FormControl
                                type="search"
                                placeholder="Search"
                                className="me-2"
                                aria-label="Search"
                            />
                            <Button variant="warning">Search</Button>
                        </Form>

                        {/* tombol Keranjang */}
                        <div class="keranjang">
                        <button
                            class="btn "
                            type="submit"
                            onClick={() => navigateTo(PAGE_CART)}
                        >
                            <BsCartFill />{" "}
                            <span class="badge bg-dark text-white ms-1 rounded-pill">
                                ({cart.length})
                            </span>
                        </button>
                        </div>
                        {/* tombol keranjang */}

                        <div className="logbut">
                            <button
                                class="btn "
                                type="submit"
                                onClick={() => navigateTo(PAGE_CART)}
                            >
                                <BsFillPersonFill />
                            </button>
                        </div>
                    </Navbar.Collapse>
                </Container>
            </Navbar>
            {/* navbar */}

            {/* carousel */}
            <HomeCarousel />
            {/* carousel */}

            <header>
                <button onClick={() => navigateTo(PAGE_PRODUCTS)}> Daftar Buku </button>
                <div>{/* {(PAGE_PRODUCTS)} */}</div>
            </header>

            {page === PAGE_PRODUCTS && <Products addToCart={addToCart} />}
            {page === PAGE_CART && (
                <Cart cart={cart} removeFromCart={removeFromCart} />
            )}

            {/* footer */}
            <footer className="bg-dark page-footer font-small blue pt-4 text-white">
                <div className="container-fluid text-center text-md-left">
                    <div className="row">
                        <div className="col-md-6 mt-md-0 mt-3">
                            <h5 className="text-uppercase">PesBuk</h5>
                            <p>
                                PesBuk merupakan singkatan dari Pesan Buku, website ini
                                bertujuan unutuk penjualan buku secara online dari berbagai
                                macam kategori, yang dapat dijangkau dan diakses dimanapun
                                unutuk mempermudah customer dalam memenuhi keinginan, atau
                                hobbynya.
                            </p>
                        </div>

                        <hr className="clearfix w-100 d-md-none pb-0" />

                        <div className="col-md-3 mb-md-0 mb-3">
                            <h5 className="text-uppercase">Lainnya</h5>
                            <ul className="list-unstyled">
                                <li>
                                    <a href="#!">Sayrat & Ketentuan</a>
                                </li>
                                <li>
                                    <a href="#!">Kebijakan & Privasi</a>
                                </li>
                                <li>
                                    <a href="#!">Bantuan</a>
                                </li>
                                <li>
                                    <a href="#!">Hubungi Kami</a>
                                </li>
                            </ul>
                        </div>

                        <div className="col-md-3 mb-md-0 mb-3">
                            <h5 className="text-uppercase">Tentang PesBuk</h5>
                            <ul className="list-unstyled">
                                <li>
                                    <a href="#!">Kontak</a>
                                </li>
                                <li>
                                    <a href="#!">Toko</a>
                                </li>
                                <li>
                                    <a href="#!">Kerjasama</a>
                                </li>
                                <li>
                                    <a href="#!">Tentang Kami</a>
                                </li>
                            </ul>
                        </div>
                    </div>
                </div>
                <div className="footer-copyright text-center py-3">
                    © 2020 Copyright:
                    <a href="https://pesbuk.com/"> www.pesbuk.com</a>
                </div>
            </footer>
        </div>
    );
};

export default Dashboard;
