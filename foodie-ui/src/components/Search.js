import React from 'react'

import Button from './Button';

const Search = ({onClicked}) => {
    return (
        <div>
            <input type = "text" placeholder = "Search Cuisine" />
            <Button color = "green" text = "search" onClick = {onClicked}/>
        </div>
    )
}

export default Search
