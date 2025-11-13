Data je niska **S** koja sadrži karaktere iz skupa zagrada {'(', ')', '{', '}', '[', ']'}.
Odrediti koliki je minimalan broj karaktera koje treba dodati u nisku tako da bude formiran validan izraz zagrada.  
(https://leetcode.com/problems/minimum-add-to-make-parentheses-valid/description/)

Primer:  
S = { ] { ] ( )

Rešenje: 4

Objašnjenje:
Zatvorili smo dve vitičaste zagrade na kraju. Takođe smo dodali odgovarajuće otvorene zagrade uglastim zagradama.  
{ **[** ] { **[** ] () **}** **}**