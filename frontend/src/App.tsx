import { useEffect, useState } from 'react'
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
  const [totalClicks, setTotalClicks] = useState(0)

  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    const fetchTotalClicks = async () => {
      try {
        const response = await fetch("http://localhost:3000/total_clicks");
        const result = await response.json();
        console.log("Total Clicks:", result.total_clicks);
        setTotalClicks(result.total_clicks);
      } catch (error) {
        console.error("Error fetching total clicks:", error);
      }
    };

    const intervalId = setInterval(fetchTotalClicks, 1000);

    return () => clearInterval(intervalId); // Cleanup on component unmount
  }, []);


  const handleClick = async () => {
    setError(null)
    setData({ user: "Parth", timestamp: Date.now() })
    const response = await fetch("http://localhost:3000/click", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify(data)
    })

    if (response.status !== 200) {
      setError(await response.text());
      return;
    }

    console.log(await response.text())


  }

  return (
    <>
      <div className="text-3xl">
        One Million Clicks
      </div>
      <button onClick={handleClick} className="border border-black rounded-md p-2">
        count is {totalClicks}
      </button>
      <div className="text-red-500">
        {error && <p>{error}</p>}
      </div >
    </>
  )
}

export default App
