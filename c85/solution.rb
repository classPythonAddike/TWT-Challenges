require'bigdecimal'
gets;$<.map{p 5*(BigDecimal(gets)*gets.split.map(&:to_i).sum/5).ceil}
