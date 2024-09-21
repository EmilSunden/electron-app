import React from 'react'

const Container = ({children}) => {
  return (
    <div className='flex flex-1 border border-collapse gap-8'>
      {children}
    </div>
  )
}

export default Container
