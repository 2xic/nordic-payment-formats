# Noric payment formats
There are some many "weird" banking payment file formats. Surprisingly these are the format that are the bedrock of the financial system we have here in Norway. Luckily we are starting to move to more decent file formats like ISO 20022.

### OCR 
OCR is a weird file format. Super inefficient, but hey it works. Here is the [system manual by Nets](https://www.nets.eu/no-nb/PublishingImages/Lists/Accordion%20%20OCR%20giro/AllItems/OCR%20giro%20-%20System%20manual.pdf) used for writing the parser. 

Test files used for testing was found in [python-netsgiro](https://github.com/otovo/python-netsgiro/tree/master/tests/data) and [java-netsgiro](https://github.com/Ondkloss/java-netsgiro/tree/2fd2a2f182e5d8731ce0b37191c5c128da4217e6/src/test/resources

)

### BgMax
BgMax is another weird file format. System manual used to implement was found [here](https://www.bankgirot.se/globalassets/dokument/tekniska-manualer/bankgiroinbetalningar_tekniskmanual_sv.pdf). The layout is more or less the same as a OCR file.

The official test files were used for testing https://www.bankgirot.se/tjanster/inbetalningar/bankgiro-inbetalningar/teknisk-information/

### Cremul
Cremul is maybe the best format if we were to compare it with BgMax and OCR.

It's more "efficient" by allow variable length, and don't use all the "filler" values like BgMax / OCR.

Used the documentation provided by [Nets](https://www.nets.eu/no-nb/SiteCollectionDocuments/Egiro/Implementation%20Guidelines%20CREMUL%20(ENG).pdf) and [Nordea](https://www.nordea.no/Images/152-69443/CREMUL-implementeringsguide.pdf) to the parser.

Test files used for testing was found in [cremul-parser](https://github.com/perspilling/cremul-parser/tree/0531ecb4a30e901e51e6f81d17a2cd764ced1f0c/test/unit/files) .
