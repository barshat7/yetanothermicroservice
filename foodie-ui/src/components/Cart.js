import React from 'react'
import {Link} from 'react-router-dom';

const Cart = ({totalItems, onCheckout}) => {
    return (
        <div className = "cart">
            <h3> Your Cart </h3>
            <p>Total Items <b>{totalItems}</b></p>
            <Link to = "/order">Checkout</Link>
        </div>
    )
}

export default Cart
