import React, { useState } from 'react';

const api_url = 'http://localhost:8080/fetch';

const MyComponent = () => {
  const [id, setID] = useState(0);
  const [Post, setPost] = useState('');
  const [fetchedData, setFetchedData] = useState([]);

  const getData = () => {
    fetch(api_url)
      .then((response) => response.json())
      .then((data) => setFetchedData(data))
      .catch((error) => console.log(error));
  };

  const createPost = () => {
    const newPost = { id: id, Post: Post };

    fetch('http://localhost:8080/add', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(newPost),
    })
      .then(() => {
        setID(0);
        setPost('');
        getData();
      })
      .catch((error) => console.error(error));
  };

  return (
    <div>
      <title>Campus Fora</title>
      <h1>Campus Fora</h1>

      <div className="form">
        <h2>GET ALL</h2>
        <div>
          <button onClick={getData}>Get all</button>
        </div>
        <h2>Fetched Data:</h2>
      <ul>
        {fetchedData.map((dataItem) => (
          <li key={dataItem['id']}>
            <strong>ID:</strong> {dataItem['id']} | {dataItem['Post']}
          </li>
        ))}
      </ul>

        <h2>POST</h2>
        <form onSubmit={createPost}>
          <input
            type="number"
            placeholder="ID"
            value={id}
            onChange={(e) => setID(parseInt(e.target.value))}
          />
          <input
            type="text"
            placeholder="Post"
            value={Post}
            onChange={(e) => setPost(e.target.value)}
          />
          <button type="submit">Create Post</button>
        </form>

        <h2>DELETE</h2>
        <div>
          <input placeholder="enter id"></input>
          <br></br>
          <button>Delete</button>
        </div>
      </div>

      
    </div>
  );
};

export default MyComponent;
