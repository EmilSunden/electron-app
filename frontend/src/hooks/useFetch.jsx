import {useState, useEffect} from 'react'

export const useFetch = () => {
    const [data, setData] = useState(null)
    const [error, setError] = useState(null)
    const [loading, setLoading] = useState(false)
    const [url, setUrl] = useState(null)
    const [options, setOptions] = useState(null)

    useEffect(() => {
        if(!url) return;

        const fetchData = async () => {
            setLoading(true);

            const controller = new AbortController();
            const signal = controller.signal;

            try {
                const response = await fetch(url, {...options, signal});
                if(!response.ok) {
                    throw new Error("Network response was not ok")
                }
                const json = await response.json()
                setData(json)
            } catch (error) {
                if (error.name != "AbortError") {
                    setError(error);
                }
            } finally {
                setLoading(false)
            }
            return () => controller.abort();
        }
        fetchData()
    }, [url, options])

    return {data, error, loading, setUrl, setOptions}
}

