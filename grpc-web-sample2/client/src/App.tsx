import React, { useState } from 'react'

import { HelloRequest } from './helloworld/helloworld_pb'
import { GreeterClient } from './helloworld/HelloworldServiceClientPb'

function App() {
  const [inputText, setInputText] = useState("World")
  const [message, setMessage] = useState("")

  return (
    <div>
      <input
        type="text"
        value={inputText}
        onChange={(e) => setInputText(e.target.value)}
      />
      <button onClick={() => {
        const request = new HelloRequest();
        request.setName(inputText)

        const client = new GreeterClient("http://localhost:8080", {}, {})
        client.sayHello(request, {}, (err, res) => {
          if (err || res === null) {
            throw err;
          }
          setMessage(res.getMessage)
        })
      }}>Send</button>
      <p>{message}</p>
    </div>
  )
}

export default App