import React from "react";
import { Route, Link } from "react-router-dom";

import Navbar from "../components/Navbar";
// import Profile from "../page/Profile";
// import Keranjang from "../page/Keranjang";

import biologyBook from "../asset/img/Item/Biology.jpeg";

import "../asset/css/Dashboard.css";

function DashboardUi () {
    return (
        <div className="Dashboard">
            {/* <h1>Dashboard</h1> */}
            <Navbar />
            {/* <Route exact path="/" component={home}></Route>
            <Route exact path="/keranjang" component={Keranjang}></Route>
            <Route exact path="/profil" component={Profile}></Route> */}

            <div className= "dashboard-container">
                <div className="display">
                    <ul>
                        <li>
                            <a href="#">
                                <img src={biologyBook} alt="biology book" width={250} height={250}/>
                                <p>buku biologi</p>
                            </a>
                        </li>
                        <li>
                            <a href="#">
                                <img src={biologyBook} alt="biology book" width={250} height={250}/>
                                <p>buku biologi</p>
                            </a>
                        </li>
                        <li>
                            <a href="#">
                                <img src={biologyBook} alt="biology book" width={250} height={250}/>
                                <p>buku biologi</p>
                            </a>
                        </li>
                        <li>
                            <a href="#">
                                <img src={biologyBook} alt="biology book" width={250} height={250}/>
                                <p>buku biologi</p>
                            </a>
                        </li>
                        <li>
                            <a href="#">
                                <img src={biologyBook} alt="biology book" width={250} height={250}/>
                                <p>buku biologi</p>
                            </a>
                        </li>
                        <li>
                            <a href="#">
                                <img src={biologyBook} alt="biology book" width={250} height={250}/>
                                <p>buku biologi</p>
                            </a>
                        </li>
                        <li>
                            <a href="#">
                                <img src={biologyBook} alt="biology book" width={250} height={250}/>
                                <p>buku biologi</p>
                            </a>
                        </li>
                        <li>
                            <a href="#">
                                <img src={biologyBook} alt="biology book" width={250} height={250}/>
                                <p>buku biologi</p>
                            </a>
                        </li>
                        <li>
                            <a href="#">
                                <img src={biologyBook} alt="biology book" width={250} height={250}/>
                                <p>buku biologi</p>
                            </a>
                        </li>
                        <li>
                            <a href="#">
                                <img src={biologyBook} alt="biology book" width={250} height={250}/>
                                <p>buku biologi</p>
                            </a>
                        </li>
                    </ul>
                </div>     
            </div>
        </div>
    );
}

export default DashboardUi;