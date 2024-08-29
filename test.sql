select userID
      ,timestamp
      ,anonymousId
      ,geo
  from `{{ params.web_project }}.mytable.Users`
 where geo = 'IE'
