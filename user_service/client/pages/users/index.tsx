import { useEffect, useState } from 'react'
import { NextPage } from 'next'
import { User } from 'proto/user_pb'
import { UserServiceClient } from 'proto/UserServiceClientPb'
import { Empty } from 'google-protobuf/google/protobuf/empty_pb'
import { apiEndpoint } from 'resources/constants'

const UsersIndexPage: NextPage = () => {
  const [users, setUsers] = useState<Array<User>>([])

  useEffect(() => {
    const userServiceClient = new UserServiceClient(apiEndpoint)
    userServiceClient.getUsers(new Empty(), {}, (_, res) => {
      const usersList = res.getUsersList()
      setUsers(usersList)
    })
  }, [])

  return (
    <>
      <h1>ユーザー一覧</h1>
      {users.map((user, index) => {
        return (
          <div key={index}>
            <div>
              <a href={`/users/${user.getName()}`}>{user.getName()}</a>
            </div>
          </div>
        )
      })}
    </>
  )
}

export default UsersIndexPage
