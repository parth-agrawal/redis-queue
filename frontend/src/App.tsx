import { useEffect, useState } from 'react'
import './App.css'

type ClickRequest = {
  user: string,
  timestamp: number
}

const BACKEND_URL = import.meta.env.VITE_BACKEND_URL || "http://localhost:3000";

function App() {
  const [data, setData] = useState<ClickRequest>({ user: "Parth", timestamp: Date.now() })
  const [totalClicks, setTotalClicks] = useState(0)

  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    const errorTimeout = setTimeout(() => {
      setError(null);
    }, 3000);

    return () => clearTimeout(errorTimeout); // Cleanup on component unmount
  }, [error]);

  useEffect(() => {
    const fetchTotalClicks = async () => {
      try {
        const response = await fetch(`${BACKEND_URL}/total_clicks`);
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
    const response = await fetch(`${BACKEND_URL}/click`, {
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
      <div className="flex flex-col justify-center items-center gap-5 h-screen w-screen">

        <div className="absolute inset-0"
          style={{ backgroundImage: `url(pointer.png)`, backgroundSize: '20px 20px', backgroundRepeat: 'repeat', opacity: 0.1 }}>
        </div>
        <div className="relative text-3xl" style={{ opacity: 1 }} >
          One Million Clicks
        </div>
        <button onClick={handleClick} className="relative border border-black rounded-md p-2" style={{ opacity: 1 }}>
          count is {totalClicks} / 1,000,000
        </button>
        <div className="relative text-red-500 min-h-[20px]" style={{ opacity: 1 }}>
          {error ? <p>{error}</p> : <p>&nbsp;</p>}
        </div>

      </div>
    </>
  )
}

export default App
