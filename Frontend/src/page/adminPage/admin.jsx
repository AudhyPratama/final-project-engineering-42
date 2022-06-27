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
// import DashboardUi from './page/dashboard/Dashboard';

// import react icon
import { BsCartFill } from "react-icons/bs";
import { BsFillPersonFill } from "react-icons/bs";
import HomeCarousel from "../keranjang/HomeCarousel";




// import css
import "./admin.css";

const Admin = () => {
    const navigate = useNavigate();
    
    const [book_name, setBookName] = useState('');
    const [categori_id, setCategoriId] = useState('');
    const [penulis, setPenulis] = useState('');
    const [penerbit, setPenerbit] = useState('');
    const [kondisi, setKondisi] = useState('');
    const [berat, setBerat] = useState('');
    const [stock, setStock] = useState('');
    const [harga, setHarga] = useState('');
    const [deskripsi, setDeskripsi] = useState('');
    const [image, setImage] = useState('');

    const handleSubmit = (e) => {
        e.preventDefault();
        axios.post('http://localhost:8080/api/admin/product/add', {
            categori_id,    
            book_name,
            penulis,
            penerbit,
            kondisi,
            stock,
            berat,
            harga,
            deskripsi,
            image
        }).then(res => {
            console.log(res);
            console.log(res.data);
            alert('Berhasil Menambahkan Data');
        }
        ).catch(err => {
            console.log(err);
        });
    };

    
    const handleUpdate = (e) => {
        e.preventDefault();
        axios.post('http://localhost:8080/api/admin/product/update', {
            categori_id,
            book_name,
            penulis,
            penerbit,
            kondisi,
            stock,
            berat,
            harga,
            deskripsi,
            image
        }).then(res => {
            console.log(res);
            console.log(res.data);
            alert('Berhasil Mengupdate Data');
        }
        ).catch(err => {
            console.log(err);
        });
    };


    

    return (
        <div className="Dashboard">
            {/* navbar */}
            <Navbar bg="light" expand="lg">
                <Container>
                    <Navbar.Brand href="#admin">Admin</Navbar.Brand>
                    <Navbar.Toggle aria-controls="basic-navbar-nav" />
                    <Navbar.Collapse id="basic-navbar-nav">
                        <div className="logbut">
                            <button
                                class="btn "
                                type="submit"
                                href="/"
                                onClick={() => navigate('/')}
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
                <div className="container">
                    <div className="row">
                        <div className="col-md-12">
                            <div className="card">
                                <div className="card-header">
                                    <h4>Buku</h4>
                                </div>
                                <div className="card-body">
                                    <div className="table-responsive">
                                        <table className="table table-striped">
                                            <thead>
                                                <tr>
                                                    <th>No</th>
                                                    <th>CategoryID</th>
                                                    <th>JudulBuku</th>
                                                    <th>Penulis</th>
                                                    <th>Penerbit</th>
                                                    <th>Kondisi</th>
                                                    <th>Berat</th>
                                                    <th>Stok</th>
                                                    <th>Harga</th>
                                                    <th>Deskripsi</th>
                                                    <th>Gambar</th>
                                                    <th>Aksi</th>
                                                </tr>
                                            </thead>
                                            <tbody>
                                                <tr>
                                                    <td>1</td>
                                                    <td>
                                                        <input
                                                            type="text"
                                                            className="form-control"
                                                            value={categori_id}
                                                            // cara menyimpan category_id menjadi integer
                                                            onChange={(e) => setCategoriId(parseInt(e.target.value))}
                                                        />
                                                    </td>
                                                    <td>
                                                        <input
                                                            type="text"
                                                            className="form-control"
                                                            placeholder="Judul Buku"
                                                            value={book_name}
                                                            onChange={(e) => setBookName(e.target.value)}
                                                        />
                                                    </td>
                                                    <td>
                                                        <input
                                                            type="text"
                                                            className="form-control"
                                                            placeholder="Penulis"
                                                            value={penulis}
                                                            onChange={(e) => setPenulis(e.target.value)}
                                                        />
                                                    </td>
                                                    <td>
                                                        <input
                                                            type="text"
                                                            className="form-control"
                                                            placeholder="Penerbit"
                                                            value={penerbit}
                                                            onChange={(e) => setPenerbit(e.target.value)}
                                                        />
                                                    </td>
                                                    <td>
                                                        <input
                                                            type="text"
                                                            className="form-control"
                                                            placeholder="Kondisi"
                                                            value={kondisi}
                                                            onChange={(e) => setKondisi(e.target.value)}
                                                        />
                                                    </td>
                                                    <td>
                                                        <input
                                                            type="text"
                                                            className="form-control"
                                                            placeholder="Berat"
                                                            value={berat}
                                                            onChange={(e) => setBerat(e.target.value)}
                                                        />
                                                    </td>
                                                    <td>
                                                        <input
                                                            type="number"
                                                            className="form-control"
                                                            
                                                            // cara menyimpan stock menjadi integer
                                                            value={stock}
                                                            onChange={(e) => setStock(parseInt(e.target.value))}
                                                            //onChange={(e) => setStock(e.target.value)}
                                                        />
                                                    </td>
                                                    <td>
                                                        <input
                                                            type="text"
                                                            className="form-control"
                                                            placeholder="Harga"
                                                            value={harga}
                                                            // cara menyimpan harga menjadi integer
                                                            onChange={(e) => setHarga(parseInt(e.target.value))}
                                                        />
                                                    </td>
                                                    <td>
                                                        <input
                                                            type="text"
                                                            className="form-control"
                                                            placeholder="Deskripsi"
                                                            value={deskripsi}
                                                            onChange={(e) => setDeskripsi(e.target.value)}
                                                        />
                                                    </td>
                                                    <td>
                                                        <input
                                                            type="text"
                                                            className="form-control"
                                                            placeholder="Gambar"
                                                            value={image}
                                                            onChange={(e) => setImage(e.target.value)}
                                                        />
                                                    </td>
                                                    <td>
                                                        <button
                                                            className="btn btn-primary"
                                                            type="submit"
                                                            onClick={handleSubmit}
                                                        >
                                                            Simpan  
                                                        </button>
                                                    </td>
                                                </tr>
                                            </tbody>
                                        </table>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
              

            </header>

            
        </div>
    );
};

export default Admin;