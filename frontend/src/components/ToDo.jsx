import { useEffect, useRef } from "react"
import { gsap } from "gsap";

import {removeToDo, updateToDo} from '../services/ToDo'

export default function({ID, Title, Done, removeToDo, setReload}){
    const ToDoElement = useRef(null)
    const checkboxElement = useRef(null)
    let animationDuration = 250

    useEffect(() => {
        gsap.fromTo(ToDoElement.current, {opacity: .25}, {opacity: 1, duration: animationDuration/1000})
    }, [])

    function removeToDoHandler() {
        checkboxElement.current.setAttribute('disabled', '')

        gsap.to(ToDoElement.current, {opacity: 0, duration: animationDuration/1000})

        setTimeout(() => {
            removeToDo(ID, setReload)
        }, animationDuration)
    }

    return (
        <li key={ID} ref={ToDoElement} className='flex justify-start items-center gap-2 border-b border-gray-300 pb-1'>
            <div>
            <input ref={checkboxElement} onChange={removeToDoHandler} autoComplete='false' spellCheck="false" type="checkbox" name="" id="" />
            </div>
            <input onChange={e => updateToDo(ID, e.target.value, setReload)} type="text" className="focus:outline-none" autoComplete='false' spellCheck='false' defaultValue={Title}/>
        </li>
    )
}