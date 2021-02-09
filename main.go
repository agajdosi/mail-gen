package main

import (
	"fmt"

	"github.com/lucasjones/reggen"
)

func main() {

	template := "(Dobrý den|Zdravím Vás),\n\n" +

		"(při reinstalaci|během repase|při opravě|během generální údržby|při odvirování|po odvirování|při forenzní analýze|během forenzní analýzy) " +
		"(jednoho ze serverů|hlavního serveru|datacentra|diskového pole) " +
		"(našeho zákazníka|naší zákaznice|našeho uživatele|naší uživatelky|významného zákazníka|zákazníka globálního významu) " +
		"(se nám podařilo získat|jsme narazili na|se nám podařilo obnovit|jsme z poškozeného disku obnovili|jsme na disku objevili)" +
		", (patrně|nejspíše|s nejvyšší pravděpodobností|podle všeho) " +
		"(zcizen|odcizen|ztracen|zapomenut|hacknut)(ou, obrazovou dokumentaci|é, (fotografie|obrázkové soubory)|á, obrazová data) " +
		"(Vašich (uměleckých|výtvarných) (prací|děl)|Vaší (umělecké|výtvarné) (tvorby|výroby))" +
		", jenž (shodou okolností|díky bohu|naštěstí|zaplaťpánbů) (na první pohled|ihned|téměř okamžitě|bryskně) " +
		"(rozpoznal|identifikoval)(a (jedna z mých kolegyň|moje vedoucí|jedna naše internistka|jedna z našich brigádnic)| (jeden z mých kolegů|můj vedoucí|jeden náš internista|jeden z našich brigádníků))\\." +

		" Vzhledem k (GDPR|platné (evropské |české |)legislativě) i (potenciálně|relativně|nejspíše) (vysoké historické|umělecké|kunsthistorické|osobní|archivní) hodnotě" +
		" (bylo|je) (naší|mojí) \\(nejen zákonnou\\) povinností Vás o tomto (výjimečném|překvapivém|dechberoucím) nálezu (informovat|vyrozumět)\\. " +

		"A to (zejména|hlavně|především) (kvůli|vzhledem) k (tomu|faktu), že (nalezené|objevené|obnovené) (soubory|filey) " +
		"jsou (ve vyšší kvalitě|mnohem kvalitnější|ve vyšším rozlišení|řádově kvalitnější|řádově obrazově lepší) než ty, " +
		"které (užíváte|máte|prezentujete|publikujete) na Vašem (webovém portfoliu|online portfoliu|uměleckém webu) " +
		"- a tedy (mohou|by mohla) (potenciálně|případně) posloužit k (dodatečnému|dalšímu) (vy|z)lepšení Vaší (prezentace|webové stránky)\\.\n\n" +

		"Dejte (mi|nám), prosím, vědět zda máte zájem o zaslání (nalezených|obnovených|inkriminovaných) (souborů|obrázků), " +
		"(nebo|či) zda naopak (upřednostníte jejich smazání|dáváte souhlas k jejich odstranění|souhlasíte s jejich smazáním)\\.\n\n" +

		"(S pozdravy|Se srdečnými pozdravy|S přáním pěkného dne),\nIng\\. Jan Štolfa, MBA"

	generator, err := reggen.NewGenerator(template)
	if err != nil {
		return
	}

	text := generator.Generate(10)

	fmt.Println(text)
}
