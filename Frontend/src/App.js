import React, { useState } from "react";

import "./App.css";

import Products from "../src/page/keranjang/Products";
import Cart from "../src/page/keranjang/Cart";

import DashboardUi from './page/dashboard/Dashboard';
// import LoginUi from "./page/login/Login";


// keranjang belanja
const PAGE_PRODUCTS = 'products';
const PAGE_CART = 'cart';
// keranjang belanja
  


function App() {
  // keranjang belanja
  const [cart, setCart] = useState([]);
  const [page, setPage] = useState([PAGE_PRODUCTS]);


  const removeFromCart = (productToRemove) => {
    setCart(cart.filter((product) => product !== productToRemove));
  };

  const addToCart = (product) => {
    console.log('add cart');
    setCart([...cart, {...product}]);
  };

  const navigateTo = (nextPage) => {
    setPage(nextPage);
  }
  // keranjang belanja

  return (
    <div className='App'>
      {/* <LoginUi /> */}
      <DashboardUi />
      <header>
        <button onClick={() => navigateTo(PAGE_CART)}>Keranjang ({cart.length})</button>
        <button onClick={() => navigateTo(PAGE_PRODUCTS)}> Daftar Buku </button>
      </header>
      {page === PAGE_PRODUCTS && (<Products addToCart={addToCart} />)}
      {page === PAGE_CART && (<Cart cart = {cart} removeFromCart = {removeFromCart} />)}
    </div> 
  );
}

export default App;
