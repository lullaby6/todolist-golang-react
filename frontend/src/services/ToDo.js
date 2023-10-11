// import { v4 as uuidv4 } from "uuid";
const API_URL = import.meta.env.VITE_API_URL

export async function addToDo(event, setReload){
    event.preventDefault();
    const formData = new FormData(event.target);
    const formProps = Object.fromEntries(formData);
    event.target.querySelectorAll('input').forEach(event => {event.value = ''})

    // setTodos([...todos, {...formProps, id: uuidv4()}])

    await fetch(`${API_URL}/api/todo`, {
      method: 'POST',
      body: JSON.stringify(formProps)
    })

    setReload(true)
}

export async function updateToDo(id, title, setReload) {
    await fetch(`${API_URL}/api/todo/${id}`, {
      method: 'PUT',
      body: JSON.stringify({
        title
      })
    })

    setReload(true)
}

export async function removeToDo(id, setReload){
    // setTodos(todos.filter(todo => todo.id !== id))
    await fetch(`${API_URL}/api/todo/${id}`, {method: 'DELETE'})
    setReload(true)
}