import React from "react";

import { Link } from "react-router-dom";

import cartIcon from "../asset/icons/shopping-cart-2-128p.png";
import userIcon from "../asset/icons/user.png";
import searchIcon from '../asset/icons/search-white.png';

import '../asset/css/Navbar.css';

function Navbar () {
    return (
        <div className="page">
            <div className="Navbar">
                <ul>
                    <li>
                        {/* <Link to= "/"><h1 className="title-1">Pes<span>Buk</span></h1></Link> */}
                        <a className="App-logo" href="/home"><h2>Pes<span>Buk</span></h2></a>
                    </li>
                    
                    <li className="search-box"> 
                        <input className="input-field" type="text" name="" placeholder="Cari: Buku SMP"></input>
                    </li>

                    <li>
                        <a className="search-button" href="#"><img className="search-icon" src={searchIcon} alt="search icon" width={50} height={50}/></a>
                    </li>
                    
                    <li>
                        {/* <Link to="/keranjang">Keranjang</Link> */}
                        <a href="/keranjang"><img className="cart-icon" src={cartIcon} alt="cart icon" width={50} height={50}></img></a>
                    </li>

                    <li>
                        <p>Ini pembatas keranjang dengan profil</p>
                    </li>

                    <li>
                        {/* <Link to="/profil">Profile</Link> */}
                        <a href="/profil"><img className="user-icon" src={userIcon} alt="user icon" width={65} height={65}></img></a>
                    </li>
                </ul>
            </div>
        </div>
    );
}

export default Navbar;