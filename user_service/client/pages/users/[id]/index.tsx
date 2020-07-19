import { NextPage } from 'next'
import { useRouter } from 'next/router'

const User: NextPage = () => {
  const { id } = useRouter().query

  return <p>User ID: {id}</p>
}

export default User
