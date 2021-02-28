import React from 'react'
import Item from './Item';

const SearchResult = ({items, onAdd}) => {
    return (
        <div>
            {items.map(item => (
                <Item key = {item.id} item = {item} onAdd = {onAdd}/>
            ))}
        </div>
    )
}

export default SearchResult
