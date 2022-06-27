// import module
import axios from "axios";
import { useState } from "react";
import { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import "bootstrap/dist/css/bootstrap.min.css";
import {
    Navbar,
    Container,
    Nav,
    Button,
    Form,
    FormControl,
} from "react-bootstrap";


// import file

// import css
import "./payment.css";

const Payment = () => {

    const navigate = useNavigate();

    const cancelOrder = () => {
        alert("Order Cancelled");
        navigate("/dashboard");
    };

   
    const confirmOrder = () => {
        alert("Order Confirmed");
        navigate("/dashboard");
    };
    
    

    return (
        <div className="Dashboard">
            {/* navbar */}
            <Navbar bg="light" expand="lg">
                
                    <Navbar.Brand href="#home">Checkout</Navbar.Brand>
                    
            </Navbar>

            <header>
                <div className="container">
                    <div className="row">
                        <div className="col-md-12">
                            <div className="header-content">
                                <div className="header-content-inner">
                                    <h2>Pesanan anda akan segera kami proses</h2>
                                    <hr />

                                    <div className="row">
                                        <div className="col-md-12">
                                            <div className="form-group">
                                                <label htmlFor="">Nama</label>
                                                <input type="text" className="form-control" id="name" placeholder="Nama" />
                                            </div>
                                        </div>
                                    </div>
                                    <div className="row">
                                        <div className="col-md-12">
                                            <div className="form-group">
                                                <label htmlFor="">Alamat</label>
                                                <input type="text" className="form-control" id="alamat" placeholder="Alamat" />
                                            </div>
                                        </div>
                                    </div>
                                    <div className="row">
                                        <div className="col-md-12">
                                            <div className="form-group">
                                                <label htmlFor="">No. Telp</label>
                                                <input type="text" className="form-control" id="no_telp" placeholder="No. Telp" />
                                            </div>
                                        </div>
                                    </div>
                                    <div className="row">
                                        <div className="col-md-12">
                                            <div className="form-group">
                                                <label htmlFor="">Email</label>
                                                <input type="text" className="form-control" id="email" placeholder="Email" />
                                            </div>
                                        </div>
                                    </div>
                                    <div className="row">
                                        <div className="col-md-12">
                                            <div className="form-group">
                                                <label htmlFor="">Kode Pos</label>
                                                <input type="text" className="form-control" id="kode_pos" placeholder="Kode Pos" />
                                            </div>
                                        </div>
                                    </div>
                                    <div className="row">
                                        <div className="col-md-12">
                                            <div className="form-group">
                                                <label htmlFor="">Pilih Metode Pembayaran</label>
                                                <select className="form-control" id="metode_pembayaran">
                                                    <option value="">Pilih Metode Pembayaran</option>
                                                    <option value="1">Transfer</option>
                                                    <option value="2">COD</option>
                                                </select>
                                            </div>
                                        </div>
                                    </div>
                                    <div className="row">
                                        <div className="col-md-12">
                                            <div className="form-group">
                                                <label htmlFor="">Pilih Metode Pengiriman</label>
                                                <select className="form-control" id="metode_pengiriman">
                                                    <option value="">Pilih Metode Pengiriman</option>
                                                    <option value="1">JNE</option>
                                                    <option value="2">J&T</option>
                                                    </select>
                                            </div>
                                        </div>
                                    </div>
                                    <button type="submit" className="btn btn-primary" onClick={confirmOrder}>
                                        Bayar
                                    </button>

                                    <button type="submit" className="btn btn-danger" onClick={cancelOrder}>
                                        Batal
                                    </button>

                                </div>
                            </div>
                        </div>
                    </div>
                </div>


            </header>

        </div>
    );
};

export default Payment;