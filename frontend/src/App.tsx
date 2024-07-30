import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'

function App() {
  const [count, setCount] = useState(0)

  const handleClick = async () => {
    const response = await fetch("http://localhost:3000/click", {
      method: "POST",

    })

    console.log(response)


  }

  return (
    <>
      <div className="text-3xl">
        One Million Clicks
      </div>
      <button onClick={handleClick}>
        count is {count}
      </button>
    </>
  )
}

export default App
