import { useFetch } from "../../hooks/useFetch";

import React, { useEffect } from 'react'

import {Link} from "react-router-dom"

const ReadAllApps = () => {
    const {data, error, loading, setUrl, setOptions} = useFetch()
    
    useEffect(() => {
      setUrl('http://localhost:8080/api/get/apps');
    }, []);

    if (loading) return <div>Loading...</div>
    if (error) return <div>Error: {error.message}</div>

    console.log(data)
    console.log(error)
    console.log(loading)
  return (
    <div>
      <ul>
        {data && data.map((app, i) => (
          <div key={i}>
           <a href={app.url}>
            <img src={app.logo} alt="" width="100" height="100"/>
           </a> 
          </div>
        ))}
      </ul>
    </div>
  )
}

export default ReadAllApps
