import { NextPage } from 'next'
import Router from 'next/router'
import { useCallback, useEffect, useState } from 'react'
import { User, GetUserRequest } from 'proto/user_pb'
import { UserServiceClient } from 'proto/UserServiceClientPb'
import { apiEndpoint } from 'resources/constants'

type Props = {
  name: string
}
const UserEdit: NextPage<Props> = ({ name }) => {
  const [twitterHashTag, setTwitterHashTag] = useState('')
  const [twitchChannel, setTwitchChannel] = useState('')

  useEffect(() => {
    const userServiceClient = new UserServiceClient(apiEndpoint)
    const request = new GetUserRequest()
    request.setName(name)
    userServiceClient.getUser(request, {}, (err, res) => {
      if (err) {
        return
      }

      const user = res.getUser()
      setTwitterHashTag(user.getTwitterhashtag())
      setTwitchChannel(user.getTwitchchannel())
    })
  }, [])

  const submitUpdate = useCallback(() => {
    Router.push('/users/[name]', `/users/${name}`)
  }, [])

  return (
    <>
      <h1>設定編集</h1>
      <div>
        <label>Twitterハッシュタグ</label>
        <input
          type="text"
          placeholder="#某isNight"
          value={twitterHashTag}
          onChange={(e) => setTwitterHashTag(e.target.value)}
        />
      </div>

      <div>
        <label>Twitchチャンネル</label>
        <input
          type="text"
          placeholder="#bou_is_twitch"
          value={twitchChannel}
          onChange={(e) => setTwitchChannel(e.target.value)}
        />
      </div>

      <button onClick={submitUpdate}>更新する</button>

      <div>
        <a href={`/users/${name}`}>ユーザー詳細</a>
      </div>
    </>
  )
}

UserEdit.getInitialProps = async ({ query }) => {
  return {
    name: query.name as string,
  }
}

export default UserEdit
