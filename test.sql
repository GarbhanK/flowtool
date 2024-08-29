select userID
      ,timestamp
      ,anonymousId
      ,geo
  from `{{ temptest }}.mytable.Users`
 where geo = 'IE'
