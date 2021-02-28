import React from 'react'

import Header from './Header';
import Search from './Search';
import SearchResult from './SearchResult';
import Cart from './Cart';
import {BrowserRouter as Router, Route} from 'react-router-dom';
import Order from './Order';
import { useState } from 'react';

const dbItems = [
    {
      id: 1,
      name: "Chilli Chicken",
      description: "A south-east Asian staple of crispy batter friend chicken chunks tossed in a spicy mix of garlic, chilli, onion and peppers.",
      price: 525
    },
    {
      id: 2,
      name: "Hot & Spicy Chicken Wings",
      description: "Delicious fried chicken winglets prepared in a traditional red marinade of garlic, chilli and soy sauce.",
      price: 645
    },
    {
      id: 3,
      name: "Konjee Crispy Lamb",
      description: "Crispy Shredded Lamb Coated With Hoisin Sauce & Sesame Seeds.",
      price: 579
    }
  ]


const HomePage = () => {

    const [items, setItems] = useState([]);

    const [cart, setCart] = useState([]);

    const searchClicked = () => {
        setItems(dbItems);
    }

    const addToCart = (item) => {
      let existingItem = cart.find((c) => {
        if (c.id === item.id) {
          return c;
        }
      })
      if (existingItem) {
        existingItem.quantity = existingItem.quantity + 1;
        setCart([...cart]);
      } else {
        item.quantity = 1;
        setCart([...cart, item]);
      }
    }

    const onCheckout = () => {
      console.log('Checking Out');
    }

    const onRemove = (itemId) => {
      const newCart = cart.filter((c) => c.id !== itemId);
      setCart(newCart);
    }

    const onOrder = async () => {
      await createCartAndOrder();
    }

    const createCartAndOrder = async () => {
      const res = await fetch('http://localhost:4000/carts', {
        method: 'POST',
        headers: {
          'Content-type': 'application/json'
        },
        body: JSON.stringify({
          userID: 'barshat77',
          cartItemDtoList: cart
        })
      });
      await res.json();

    }

    return (
           <Router>
              <Route 
              path = "/" exact render = {(props) => (
                <div className="main">
                <Header className= "header" />
                <Cart totalItems = {cart.length}/>
                <Search className = "search" onClicked = {searchClicked}/>
                {items.length > 0 ? <SearchResult items = {items} onAdd = {addToCart} /> : ("Search Your Favourite Dish")}
              </div>
              )
              }
              />
              <Route
                path = "/order" component = {() => <Order cartItems = {cart} onRemove = {onRemove} onOrder = {onOrder} />}
              />
           </Router>
    )
}

export default HomePage
