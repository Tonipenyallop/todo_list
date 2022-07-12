import axios from "axios";
function App() {
  const getHelloRequest = async () => {
      const helloRequest = await axios.get("/hello");
      console.log(helloRequest.data)
  }
  return (
    <div>
      <button onClick={getHelloRequest}>Click me</button>
    </div>
  )
    
  ;
}

export default App;
