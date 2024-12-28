import asyncio
import aiohttp
import time
import random

async def make_request(session, url):
    """Send a single GET request to the specified URL."""
    try:
        start_time = time.time()
        async with session.get(url + generate_random_string(3)) as response:
            duration = time.time() - start_time
            status = response.status
            data = await response.text()
            return status, duration
    except Exception as e:
        return None, None, str(e)

async def load_test(url, n):
    """Perform a load test by sending n parallel requests to the URL."""
    async with aiohttp.ClientSession() as session:
        tasks = [make_request(session, url) for _ in range(n)]
        results = await asyncio.gather(*tasks, return_exceptions=True)

    # Process results
    response_times = []
    successful_requests = 0
    failed_requests = 0

    for result in results:
        if isinstance(result, tuple):
            status, duration = result[:2]
            if status:
                successful_requests += 1
                response_times.append(duration)
            else:
                failed_requests += 1

    avg_response_time = sum(response_times) / len(response_times) if response_times else 0
    
    print("Load Test Results:")
    print(f"Total Requests: {n}")
    print(f"Successful Requests: {successful_requests}")
    print(f"Failed Requests: {failed_requests}")
    print(f"Average Response Time: {avg_response_time:.2f} seconds")

if __name__ == "__main__":
    url = "http://localhost:3000/api/v1/db/"
    n = int(input("Enter the number of parallel requests: "))

    asyncio.run(load_test(url, n))