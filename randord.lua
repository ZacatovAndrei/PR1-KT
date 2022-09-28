#! /bin/env lua
local ranks={2,1,1,3,3,1,2,2,1,1,1,2,2}
local eqdecode={"nothing","oven","stove"}
local equipment={1,0,2,0,1,2,1,1,2,0,0,0,1}
local r1freq={}
local r2freq={}
local r3freq={}
local numords={0,0,0}
math.randomseed(os.time())
local function prcalc(table)

  local opr=table.bp
  local itlen=#table
  local avgran=0.0

  for _,rank in ipairs(table) do avgran=avgran+ranks[rank] end
  avgran=avgran/itlen
  tmpres=-1+math.floor(12+6*(opr/5)-itlen-avgran)//2
  if tmpres<=0 then tmpres=0 
  elseif tmpres>=5 then tmpres=5 end
  return tmpres
end

local lenrank=function(a,b)
  if a.cp==b.cp then return #a<#b end
  return a.cp>b.cp
end

local ranklen=function(a,b)
  if #a==#b then return a.cp>b.cp end
  return #a<#b
end


local function genorders(n)
  ordtab={}
  ordtab.dishes=0
  for i=1,n do
    ordtab[i]={}
    foods=math.random(1,6)
    ordtab.dishes=ordtab.dishes+foods
    for j=1,foods do
      ordtab[i][j]=math.random(1,13);
      if ranks[ordtab[i][j]] == 1 then 
        r1freq[ordtab[i][j]]=r1freq[ordtab[i][j]]+1 
        numords[1]=numords[1]+1
      elseif ranks[ordtab[i][j]] == 2 then 
        r2freq[ordtab[i][j]]=r2freq[ordtab[i][j]]+1 
        numords[2]=numords[2]+1
      elseif ranks[ordtab[i][j]] == 3 then 
        r3freq[ordtab[i][j]]=r3freq[ordtab[i][j]]+1 
        numords[3]=numords[3]+1 
      end
    end
    ordtab[i].bp=math.random(1,5)
    ordtab[i].cp=prcalc(ordtab[i])
  end
  return ordtab
end

local function initfreqtables(ranks)
  for i,rank in ipairs(ranks) do
    if rank==1 then r1freq[i]=0
    elseif rank==2 then r2freq[i]=0
    elseif rank==3 then r3freq[i]=0 end
  end
end




-- ACTUAL START OF THE TEST PROGRAM

initfreqtables(ranks)
local orders=genorders(10)
table.sort(orders,ranklen)
for _,order in ipairs(orders) do
  io.write("["..order[1])
  for i=2,#order do io.write(","..order[i]) end
  io.write(string.format("] (%i) p=%i\n",order.bp,order.cp))
end
print(string.format(
    "There are %i dishes\n%i of rank 1 %i of rank 2 %i of rank 3\n",
    ordtab.dishes,
    numords[1],
    numords[2],
    numords[3]
  )
)


--[[
  for _,order in ipairs(orders) do
    io.write("["..order[1])
    for i=2,#order do io.write(","..order[i]) end
    io.write(string.format("] (%i) p=%i\n",order.bp,order.cp))
  end
  print(string.format("There are %i dishes\n\n",ordtab.dishes))

  table.sort(orders,ranklen)

  for _,order in ipairs(orders) do
    io.write("["..order[1])
    for i=2,#order do io.write(","..order[i]) end
    io.write(string.format("] (%i) p=%i\n",order.bp,order.cp))
  end
  print(string.format("There are %i dishes\n\n",ordtab.dishes))
  --]]