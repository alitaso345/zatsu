import { useEffect, useState } from 'react'
import { NextPage } from 'next'
import { useRouter } from 'next/router'
import { User, GetUserRequest } from 'proto/user_pb'
import { UserServiceClient } from 'proto/UserServiceClientPb'
import { apiEndpoint } from 'resources/constants'

type Props = {
  name: string
}

const UserPage: NextPage<Props> = ({ name }) => {
  const [user, setUser] = useState<User>(null)

  useEffect(() => {
    const userServiceClient = new UserServiceClient(apiEndpoint)
    const request = new GetUserRequest()
    request.setName(name)
    userServiceClient.getUser(request, {}, (err, res) => {
      if (err) {
        return
      }

      const user = res.getUser()
      setUser(user)
    })
  }, [])

  return user ? (
    <div>Name: {user.getName()}</div>
  ) : (
    <div>Not Found User {name}</div>
  )
}

UserPage.getInitialProps = async ({ query }) => {
  return {
    name: query.name as string,
  }
}

export default UserPage
