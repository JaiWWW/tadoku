export interface Contest {
  id: number
  description: string
  start: Date
  end: Date
  open: boolean
}

export interface rawContest {
  id: number
  description: string
  start: string
  end: string
  open: boolean
}
