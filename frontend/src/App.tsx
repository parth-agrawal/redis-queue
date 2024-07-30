import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'

type ClickRequest = {
  user: string,
  timestamp: number
}

function App() {
  const [count, setCount] = useState(0)
  const [data, setData] = useState<ClickRequest>()


  const handleClick = async () => {
    setData({ user: "Parth", timestamp: Date.now() })
    const response = await fetch("http://localhost:3000/click", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify(data)
    })

    console.log(await response.text())


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
