import { useState, useEffect } from 'react'
import { Icon } from '@iconify/react';
import { v4 as uuidv4 } from "uuid";

import useFetch from './hooks/useFetch';

import ToDo from './components/ToDo';

export default function App() {
  const [todos, setTodos] = useState([])
  const [reload, setReload] = useState(false)
  const {data, loading, error} = useFetch('http://localhost:3000/api/todo')

  useEffect(() => {
    if(reload) {
      setReload(false)
    }
  }, [reload])

  useEffect(() => {
    if(!loading && data) setTodos(data.data);
  })

  async function addToDo(e){
    e.preventDefault();
    const formData = new FormData(e.target);
    const formProps = Object.fromEntries(formData);
    e.target.querySelectorAll('input').forEach(e => {e.value = ''})

    // setTodos([...todos, {...formProps, id: uuidv4()}])

    await fetch(`http://localhost:3000/api/todo`, {
      method: 'POST',
      body: JSON.stringify(formProps)
    })

    setReload(true)
  }

  async function removeToDo(id){
    // setTodos(todos.filter(todo => todo.id !== id))
    await fetch(`http://localhost:3000/api/todo/${id}`, {method: 'DELETE'})
    setReload(true)
  }

  return (
    <div className='flex flex-col justify-center items-center h-screen relative'>
      <div className='text-xl'>
        <form onSubmit={addToDo} className='bg-gray-400 p-6 shadow-2xl rounded flex gap-2 relative -left-4 justify-between items-center'>
          <input autoComplete='false' required name='title' spellCheck="false" type="text" className='focus:outline-none shadow-2xl shadow-white bg-transparent text-white border-b border-b-gray-100'/>
          <button type="submit" className='text-white bg-gray-500 rounded shadow-2xl h-7 w-7 flex justify-center items-center opacity-75 transition-opacity hover:opacity-100'><Icon icon="fluent:add-12-regular" /></button>
        </form>
        <ul className='p-6 text-gray-700 shadow-2xl flex flex-col justify-center gap-2 rounded border border-gray-100 bg-white relative -top-2'>
          {todos.length > 0
            ? todos.map((todo) =>
                <ToDo key={todo.ID} {...todo} removeToDo={removeToDo}></ToDo>
              )
            : <p className='opacity-75 text-lg'>No ToDos remaining...</p>
          }
        </ul>
      </div>
    </div>
  )
}