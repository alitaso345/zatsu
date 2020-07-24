import { useEffect, useState } from 'react'
import { User } from 'proto/user_pb'
import { UserServiceClient } from 'proto/UserServiceClientPb'
import { Empty } from 'google-protobuf/google/protobuf/empty_pb'

const Users: React.FC = () => {
  const apiEndpoint =
    process.env.NODE_ENV === 'development'
      ? 'http://localhost:8080'
      : 'http://' + window.location.host
  const [users, setUsers] = useState<Array<User>>([])

  useEffect(() => {
    const userServiceClient = new UserServiceClient(apiEndpoint)
    userServiceClient.getUsers(new Empty(), {}, (_, res) => {
      const usersList = res.getUsersList()
      setUsers(usersList)
    })
  }, [])

  return (
    <div>
      <h1>Users List</h1>
      {users.map((user, index) => {
        return (
          <div key={index}>
            <div>ID: {user.getId()}</div>
            <div>Name: {user.getName()}</div>
          </div>
        )
      })}
    </div>
  )
}

export default Users
