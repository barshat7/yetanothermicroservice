import React from 'react'
import Button from './Button';

const OrderItem = ({item, onRemove}) => {
    return (
        <div className = 'order-display'>
            <h3> {item.name}  quantity: {item.quantity}</h3>
            <p>{item.price}</p>
            <Button color = 'red' text = 'remove' onClick = {() => onRemove(item.id)} />
        </div>
    )
}

export default OrderItem
