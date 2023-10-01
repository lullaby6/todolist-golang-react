import { useEffect, useRef } from "react"
import { gsap } from "gsap";

export default function({title, done, id, removeToDo}){
    const ToDoElement = useRef(null)
    const checkboxElement = useRef(null)
    let animationDuration = 750

    useEffect(() => {
        gsap.fromTo(ToDoElement.current, {opacity: .25}, {opacity: 1, duration: animationDuration/1000})
    }, [])

    function removeToDoHandler() {
        checkboxElement.current.setAttribute('disabled', '')

        gsap.to(ToDoElement.current, {opacity: 0, duration: animationDuration/1000})

        setTimeout(() => {
            removeToDo(id)
        }, animationDuration)
    }

    return (
        <li key={id} ref={ToDoElement} className='flex justify-start items-center gap-2 border-b border-gray-300 pb-1'>
            <div>
            <input ref={checkboxElement} onChange={removeToDoHandler} autoComplete='false' spellCheck="false" type="checkbox" name="" id="" />
            </div>
            <p>{title}</p>
        </li>
    )
}