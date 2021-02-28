import React from 'react'
import Button from './Button';

const Item = ({item, onAdd}) => {
    return (
        <div>
            <h3>{item.name}</h3>
            <p>{item.description}</p>
            <p>{item.price}</p>
            <Button color = 'orange' text = 'Add To Cart' onClick = {() => onAdd(item)} />
        </div>
    )
}

export default Item
