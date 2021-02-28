import React from 'react'
import OrderItem from './OrderItem';
import Button from './Button';


const Order = ({cartItems, onRemove, onOrder}) => {

    console.log("Building Cart " + JSON.stringify(cartItems));
    
    return (
        <div>
            {
            cartItems.map((item) => <OrderItem key = {item.id} item = {item} onRemove = {onRemove} />)
            }

            <Button color = "orange" text = "Place Order" onClick = {() => onOrder()}/>
        </div>
    )
}

export default Order
